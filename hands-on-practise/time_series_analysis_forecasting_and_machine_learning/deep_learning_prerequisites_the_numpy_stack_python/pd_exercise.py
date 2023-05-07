import pandas as pd
import matplotlib.pyplot as plt

from sklearn.datasets import make_circles


t = make_circles(n_samples=5000, noise=0.09, factor=0.5)

df = pd.DataFrame(t[0], columns=['x1', 'x2'])

df['x1^2'] = df['x1'] ** 2
df['x2^2'] = df['x2'] ** 2
df['x1*x2'] = df['x1'] * df['x2']

df['y'] = t[1]

plt.scatter(df['x1'], df['x2'], c=df['y'])
plt.show()

df.to_csv('make_circle_dataset.csv', index=False, header=False)