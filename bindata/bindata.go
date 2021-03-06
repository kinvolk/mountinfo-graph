package bindata

func Index() string {
	return `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>Kinvolk - Mountinfo graph generator</title> 
</head>

<body>
<h3>{{printf "%s" . }}</h3>
<form action="/show" method="post" id="miform">
    <textarea rows="30" cols="80" name="mountinfofile" form="miform"></textarea>
    <input type="submit" value="Generate"/>
</form>
</body>
</head>`
}

func Show() string {
	return `<!doctype html>
<head>
  <meta charset="utf-8">
  <title>Kinvolk - Mountinfo graph generator</title> 
</head>
<style> /* set the CSS */
body { color: #ff0000; }

.node circle {
  fill: #F2F2F2;
  stroke: #8AB3CF;
  stroke-width: 2px;
}

.node text { 
  fill: #574D4F;
  font: 12px sans-serif; 
}

.node--internal text {
  text-shadow: 0 1px 0 #fff, 0 -1px 0 #fff, 1px 0 0 #fff, -1px 0 0 #fff;
}

.link {
  fill: none;
  stroke: #969091;
  stroke-width: 1px;
}

</style>

<body>
<!-- load the d3.js library -->    	
<script src="//d3js.org/d3.v4.min.js"></script>
<script>

var treeData = JSON.parse({{.}})

// set the dimensions and margins of the diagram
var margin = {top: 100, right: 100, bottom: 100, left: 100},
    width = window.innerWidth - margin.left - margin.right,
    height = window.innerHeight - margin.top - margin.bottom;

// declares a tree layout and assigns the size
var treemap = d3.tree()
    .size([width, height]);

//  assigns the data to a hierarchy using parent-child relationships
var nodes = d3.hierarchy(treeData);

// maps the node data to the tree layout
nodes = treemap(nodes);

// append the svg obgect to the body of the page
// appends a 'group' element to 'svg'
// moves the 'group' element to the top left margin
var svg = d3.select("body").append("svg")
      .attr("width", width + margin.left + margin.right)
      .attr("height", height + margin.top + margin.bottom),
    g = svg.append("g")
      .attr("transform",
            "translate(" + margin.left + "," + margin.top + ")");

// adds the links between the nodes
var link = g.selectAll(".link")
    .data( nodes.descendants().slice(1))
  .enter().append("path")
    .attr("class", "link")
    .attr("d", function(d) {
       return "M" + d.x + "," + d.y
         + "C" + d.x + "," + (d.y + d.parent.y) / 2
         + " " + d.parent.x + "," +  (d.y + d.parent.y) / 2
         + " " + d.parent.x + "," + d.parent.y;
       });

// adds each node as a group
var node = g.selectAll(".node")
    .data(nodes.descendants())
  .enter().append("g")
    .attr("class", function(d) { 
      return "node" + 
        (d.children ? " node--internal" : " node--leaf"); })
    .attr("transform", function(d) { 
      return "translate(" + d.x + "," + d.y + ")"; });

// adds the circle to the node
node.append("circle")
  .attr("r", 10);

var str = "";

// adds the text to the node
node.append("text")
  .attr("dy", ".35em")
  .attr("y", function(d) { return d.children ? -20 : 20; })
  .style("text-anchor", "middle")
  .text(function(d) { 
                      if (d.data.name.length > 5)
                          return d.data.name.substring(d.data.name.lastIndexOf("/"), d.data.name.lastIndexOf("/")+5);
                      else
                          return d.data.name; 
        });
node.append("svg:title")
  .text(function(d) { return d.data.name });
</script>
</body>`
}
