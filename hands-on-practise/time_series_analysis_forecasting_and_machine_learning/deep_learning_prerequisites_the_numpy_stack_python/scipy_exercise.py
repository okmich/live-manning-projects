import numpy as np
import matplotlib.pyplot as plt
from scipy.signal import convolve2d
from PIL import Image


image_url = "/Users/Michael_Enudi/Downloads/IMG_20221222_094847_HDR.jpeg"

Hx = [[1, 0, -1], [2, 0, -2], [1, 0, -1]]
Hy = [[1, 2, 1], [0, 0, 0], [-1, -2, -1]]

im = Image.open(image_url)
im_arr = np.array(im)

im_arr_grayscale = np.mean(im_arr, axis=2)

# step 1: convolve Hx and Hy with grayscale image to obtain Gx and Gy
Gx = convolve2d(Hx, im_arr_grayscale)
Gy = convolve2d(Hy, im_arr_grayscale)


# step 2: take G = sqrt(Gx ^ 2 + Gy ^ 2) - this is the edge-detected output
G = np.sqrt(Gx ** 2 + Gy ** 2)

# plot new image
plt.imshow(G)
plt.show()
