"""Show map of current game for debugging purposes."""
import json
import pygame
import time

BLACK = (0, 0, 0)
WHITE = (255, 255, 255)
GREEN = (0, 255, 0)
RED = (255, 0, 0)
WIDTH = HEIGHT = 20

PLAYER_COLOR = (0, 0, 0)

def main(file_path):
    with open(file_path) as json_file:
        world_data = json.load(json_file)

    max_x_size = world_data['Max_x_size']
    max_y_size = world_data['Max_y_size']

    pygame.init()  # pylint: disable=no-member
    window_size = [max_x_size * WIDTH, max_y_size * HEIGHT]
    screen = pygame.display.set_mode(window_size)
    screen.fill(WHITE)
    pygame.display.set_caption('Game map')
    clock = pygame.time.Clock()
    while True:
        try:
            with open(file_path) as json_file:
                world_data = json.load(json_file)
        except:
            time.sleep(1)
            with open(file_path) as json_file:
                world_data = json.load(json_file)
        world_matrix = world_data['World_matrix']
        update_screen(clock, screen, world_matrix, max_x_size, max_y_size)
        # time.sleep(1)


def update_screen(clock, screen, world_matrix, max_x_size, max_y_size):
    """Update the screen of the game."""
    for x_position in range(max_x_size):
        for y_position in range(max_y_size):
            color, player = get_land_info(world_matrix, x_position, y_position)
            pygame.draw.rect(screen,
                             color,
                             [WIDTH * x_position,
                             HEIGHT * y_position,
                             WIDTH,
                             HEIGHT])

            if player:
                x_center = (x_position * WIDTH) + (WIDTH * 0.75)/2
                y_center = (y_position * HEIGHT) + (HEIGHT * 0.75)/2
                radius = (WIDTH * 0.75)/2
                pygame.draw.circle(screen,
                                   PLAYER_COLOR,
                                   (x_center, y_center),
                                   radius)

    clock.tick(60)
    pygame.display.flip()


def get_land_info(world_matrix, x_position, y_position):
    """Get info about the tile from the world matrix."""
    color = None
    terrain = world_matrix[y_position][x_position]['Terrain']
    if terrain == 'grass':
        color = (175, 255, 76)
    elif terrain == 'dirt':
        color = (155, 118, 83)
    else:
        color = (255, 255, 255)

    player = terrain = world_matrix[y_position][x_position]['Player']

    return color, player


if __name__ == "__main__":
    file_path = 'game_state.json'
    main(file_path)