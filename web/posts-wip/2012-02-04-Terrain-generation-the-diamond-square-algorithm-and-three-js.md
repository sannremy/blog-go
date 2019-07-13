https://srchea.com/terrain-generation-the-diamond-square-algorithm-and-three-js

# Terrain generation: the diamond-square algorithm and Three.js

Well, I get started learning 3D concepts/coding, so this is all new for me. I've done a first try to build an application that runs in WebGL-enabled browsers. After spending some days on it, I finally have something fit to be seen.

This post deals with how I've made that little [random terrain generation application](https://src.onl/apps/terrain-generation-diamond-square-threejs-webgl/) using the [Three.js framework](https://github.com/mrdoob/three.js/) and the [diamond-square algorithm](https://en.wikipedia.org/wiki/Diamond-square_algorithm). Obviously, this works with the new generation of web-browsers (with WebGL support). Latest Google Chrome versions, Mozilla Firefox versions and Internet Explorer 11 are supporting the WebGL context. As always, you can find the source code on my GitHub profile.

Of course, the application gives us more features than applying a texture or showing a mesh. I invite you to discover and to play with it. By the way, I would really appreciate your feedback. :-)

image

## The diamond-square algorithm

As I said above, I have used the diamond-square algorithm to generate a random fractal terrain. If you already know how the midpoint displacement algorithm works (what if you don't too), the diamond-square algorithm would be pretty easy to understand. Here is a good explanation of these algorithms.

First of all, that algorithm only works on 2D arrays of 2n+1 dimensions (e.g. 129×129, 1025×1025, etc.). As its name suggests, it works on squares and it needs the four corner points, and the midpoint to generate height values regarding average values of corners. After that, it takes the middle of each edges (it actually builds diamonds) and it takes the midpoint to get other squares (sub-squares). This excellent post by Paul Boxley shows the behavior of the algorithm step by step.

Moreover, it exists some other algorithms of terrain generations that we can mention like the Perlin noise, widely used in computer games and movies such as Tron.

## The scene

In this part, I explain how I have basically made the scene. It contains 2 main elements: a polygon mesh for the terrain and a perspective camera. Besides, it has a control element for the camera, but this will be explained in the camera section.

### Working with shapes, vertices, meshes and textures

I have used a plane (PlaneGeometry object) for the terrain grid. That shape offers the possibility to modify the z axis via their vertices.

```javascript
this.geometry = new THREE.PlaneGeometry(
    this.width,
    this.height,
    this.segments,
    this.segments
);
var index = 0;
for(var i = 0; i <= this.segments; i++) {
    for(var j = 0; j <= this.segments; j++) {
        this.geometry.vertices[index].position.z = this.terrain[i][j];
        index++;
    }
}
```
