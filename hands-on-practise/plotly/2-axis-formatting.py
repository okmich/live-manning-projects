import plotly.offline as py # using plotly offline
import plotly.graph_objs as go
import numpy as np

x = np.arange(1,11)
y1 = np.exp(x)
y2 = np.log(x)

trace0 = go.Scatter(
   x = x, y = y1, name="exponential"
)
trace1 = go.Scatter(
   x = x, y = y2, name="logarithm", yaxis='y2'
)

data = [trace0, trace1]

layout = go.Layout(
   title = "Chart with secondary Y axis", 
   yaxis = dict(title='exp', showline=True, zeroline=True),
   yaxis2 = dict(title='log', showline=True, zeroline=True, side='right', overlaying='y')
)

fig = go.Figure(data, layout)

# to save the plot to a html page and open automatically, use
fig.write_html('axis-formatting.html', auto_open=True)


# but there are other documened way of doing secondary y axis
# https://plotly.com/python/multiple-axes/