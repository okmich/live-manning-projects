import plotly.offline as py # using plotly offline
import plotly.graph_objs as go
import numpy as np
import math #needed for definition of pi

xpoints=np.arange(0, 30)
sine_points=np.sin(xpoints)
cos_points=np.cos(xpoints)

trace0 = go.Scatter(
   x = xpoints, y = sine_points
)
trace1 = go.Scatter(
   x = xpoints, y = cos_points
)
data = [trace0, trace1]

layout = go.Layout(title = "Sine & Cosine wave", xaxis = {'title':'angle'}, yaxis = {'title':'value'})

fig = go.Figure(data, layout)

# to save the plot to a html page and open automatically, use
fig.write_html('multiple-trace.html', auto_open=True)
