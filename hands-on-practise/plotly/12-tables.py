import plotly.figure_factory as ff

import pandas as pd

## data preparation
vgsales_data = pd.read_csv('datasets/vgsales.csv')
## data preparation complete

fig =  ff.create_table(vgsales_data.describe())
fig.update_layout(title_text = "Video games sales data")
fig.write_html("hello-table.html", auto_open=True)