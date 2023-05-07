from PIL import Image
import numpy as np
import matplotlib.pyplot as plt

url = "/Users/Michael_Enudi/Downloads/IMG_20221222_094847_HDR.jpeg"

im = Image.open(url)
type(im)

arr = np.array(im)

# show the pictures
plt.imshow(arr)
plt.show()

#or
plt.imshow(im)
plt.show()


#make gray scale
gray = np.mean(arr, axis=2)