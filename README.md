# percolate

A sample application showing one usecase wher the unionfind datastructure is used.

This is inspired from the Algorithms, Part I course @ http://www.coursera.org. See programming assignment 1. Note that this does not "solve" the programming assignment. Hence I dont think I am violating the honor code. Moreover this is in go and you cant submit this anyway.

Percolation. Given a composite systems comprised of randomly distributed insulating and metallic materials: what fraction of the materials need to be metallic so that the composite system is an electrical conductor? Given a porous landscape with water on the surface (or oil below), under what conditions will the water be able to drain through to the bottom (or the oil to gush through to the surface)? Scientists have defined an abstract process known as percolation to model such situations.

Model. We model a percolation system using an N-by-N grid of sites. Each site is either open or blocked. A full site is an open site that can be connected to an open site in the top row via a chain of neighboring (left, right, up, down) open sites. We say the system percolates if there is a full site in the bottom row. In other words, a system percolates if we fill all open sites connected to the top row and that process fills some open site on the bottom row. (For the insulating/metallic materials example, the open sites correspond to metallic materials, so that a system that percolates has a metallic path from top to bottom, with full sites conducting. For the porous substance example, the open sites correspond to empty space through which water might flow, so that a system that percolates lets water fill open sites, flowing from top to bottom.)

Usage : percolation -input <INPUT FILE NAME>

Each input file starts with the number of grids followed by x y pairs of sites to be opened.

For example
4
1 1
2 1
3 1
4 1

Output :
true - Indicates that the system percolates

* - - - 
* - - - 
* - - - 
* - - - 

* -> Indicates an open site

