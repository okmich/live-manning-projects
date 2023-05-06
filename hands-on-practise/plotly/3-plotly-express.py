# Plotly Express is the easy-to-use, high-level interface to Plotly, which operates on a variety of types of data and produces easy-to-style figures.
# According to the documentation, Plotly Express is a built-in part of the plotly library, and is the recommended starting point for creating most common figures. 
# It is plotly at a higher-level.
# Every Plotly Express function uses graph objects internally and returns a plotly.graph_objects.Figure instance.
# Plotly Express provides more than 30 functions for creating different types of figures. The API for these functions was carefully designed to be as consistent and easy to learn as possible, making it easy to switch from a scatter plot to a bar chart to a histogram to a sunburst chart throughout a data exploration session. 

# Plotly Express includes the following functions categorized as 

# Basics				: scatter, line, area, bar, funnel, timeline
# Part-of-Whole			: pie, sunburst, treemap, icicle, funnel_area
# 1D Distributions		: histogram, box, violin, strip, ecdf
# 2D Distributions		: density_heatmap, density_contour
# Matrix or Image Input	: imshow
# 3-Dimensional			: scatter_3d, line_3d
# Multidimensional		: scatter_matrix, parallel_coordinates, parallel_categories
# Tile Maps				: scatter_mapbox, line_mapbox, choropleth_mapbox, density_mapbox
# Outline Maps			: scatter_geo, line_geo, choropleth
# Polar Charts			: scatter_polar, line_polar, bar_polar
# Ternary Charts		: scatter_ternary, line_ternary

# See more at https://plotly.com/python/plotly-express/

# plotly express make the plotly python universe accessible through one conduit

# colors and schemes 	=> px.color
# built in datasets	 	=> px.data
# trendliens			=> px.scatter support for trendliens.

import plotly.express as px
fig = px.bar(x=["a", "b", "c"], y=[1, 3, 2])
# save but don't automatically open
fig.write_html('plotly-express.html', auto_open=False)
# start up a simple server and server page
fig.show()

# going forward, I will use only write_html with auto_open=True.