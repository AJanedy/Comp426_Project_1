Project Completed by Andrew Janedy
10/2023

Project meets all requirements of project 1 with a few caveats

The project is written in go and uses the ebitengine library
The project implements an infinite scroll with no interruptions
The project keeps and displays a score

Input to move up and down (with up and down keys) using up and down arrows
Input to shoot using the space bar

Enemies spawn randomly, there are no more than 5 on the screen at any time
Enemies that make it across the screen reduce the score and are removed and respawned
Enemies that are hit increase the score and are removed and respawned

The shot that hits the enemy does disappear BUT, I could not figure out how to remove the shot from the list without
creating an index out of bounds error when I changed the size of the list, so I just moved it off-screen.  If you 
are able to provide some insight as to how to make this happen that would be splendid

For the same reason I did not code the lasers to be removed when they reach the end of the screen

The game does play a sound when the player fires a shot and when the shot collides with an enemy