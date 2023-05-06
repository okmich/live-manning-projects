# Plotly
# 	- Figure
# 		- Data
# 			- Traces
# 			- ...  
# 		- Layout
# 			- layout options

# the main package for using plotly include

# import plotly.plotly as py - contains functionalities for interfacing between your application and the plotly server (in the case of online mode) or between the local machine and plotly (in offline mode)
# import plotly.graph_object as go - contains all classes and functionalities for the objects that make up the visible plots 
# 								 - The defined objects in this package include
# 								 - Figure
# 								 - Data
# 								 - Layout
# 								 - Traces (like Scatter, Box, Histogram, etc)	
# 								 - All graph objects are dictionary or list-like in nature.
# import plotly.tools as tools - Contains various helpful functions for facilitating and enhancing plotly experience like functions for subplot generation, embedding plots in notebooks, etc.


# Figure - represents a plot and all it holds or should hold.
# 		  e.g. fig = go.Figure(data, layout, frames)

# Data 	- defines a list of objects that are traces.
#           e.g. data = [trace1]
# Traces 	- a collectioin of data to be plotted. The type of trace objects determines what type of infograph that will be displayed. Examples include scatter, pie, box, Histogram, pie, etc.
# 		  e.g. trace1 = go.Scatter(x=xs, y=ys)
# Layout 	- defines the appearance of the plot, and plot features which are unrelated to the data. So we will be able to change things like the title, axis titles, annotations, legends, spacing, font and even draw shapes on top of your plot
# 		  e.g. layout = go.Layout(title = "Sine wave", xaxis = {'title':'angle'}, yaxis = {'title':'sine'})

# py.iplot(fig)
import plotly.offline as py # using plotly offline
import plotly.graph_objs as go
import numpy as np
import math #needed for definition of pi

xpoints=np.arange(0, math.pi*2, 0.05)
ypoints=np.sin(xpoints)

trace0 = go.Scatter(
   x = xpoints, y = ypoints
)
data = [trace0]
layout = go.Layout(title = "Sine wave", xaxis = {'title':'angle'}, yaxis = {'title':'sine'})

fig = go.Figure(data, layout)

# to save the plot to a html page and open automatically, use
fig.write_html('hello_plotly.html', auto_open=True)

# to open the plot without saving first (the page still gets created though)
py.plot(fig)

# to open the plot in a jupyter notebook
# py.iplot(fig)