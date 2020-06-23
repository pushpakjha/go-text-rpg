"""Show map of current game for debugging purposes."""
import json
import pygame
import time

BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
GREEN = (0, 255, 0)
RED = (255, 0, 0)
WIDTH = 20
HEIGHT = 20


def main(file_path):
    with open(file_path) as json_file:
        world_data = json.load(json_file)

    max_x_size = world_data['Max_x_size']
    max_y_size = world_data['Max_y_size']

    pygame.init()  # pylint: disable=no-member
    window_size = [max_x_size * WIDTH, max_y_size * HEIGHT]
    screen = pygame.display.set_mode(window_size)
    screen.fill(WHITE)
    pygame.display.set_caption('Map')
    clock = pygame.time.Clock()
    while True:
        with open(file_path) as json_file:
            world_data = json.load(json_file)
        world_matrix = world_data['World_matrix']
        update_screen(clock, screen, world_matrix, max_x_size, max_y_size)
        time.sleep(1)


def update_screen(clock, screen, world_matrix, max_x_size, max_y_size):
    """Update the screen of the game."""
    for x_position in range(max_x_size):
        for y_position in range(max_y_size):
            color = get_land_color(world_matrix, x_position, y_position)
            pygame.draw.rect(screen,
                             color,
                             [WIDTH * x_position,
                             HEIGHT * y_position,
                             WIDTH,
                             HEIGHT])
    clock.tick(60)
    pygame.display.flip()


def get_land_color(world_matrix, x_position, y_position):
    """Get the RGB color from the world matrix."""
    color = None
    terrain = world_matrix[y_position][x_position]['Terrain']
    if terrain == 'grass':
        color = (38, 50, 22)
    elif terrain == 'dirt':
        color = (155, 118, 83)
    else:
        color = (255, 255, 255)

    return color


if __name__ == "__main__":
    file_path = 'game_state.json'
    main(file_path)