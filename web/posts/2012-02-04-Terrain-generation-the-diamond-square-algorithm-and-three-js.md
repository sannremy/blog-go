# Terrain generation: the diamond-square algorithm and Three.js

Well, I get started learning 3D concepts/coding, so this is all new for me. I've done a first try to build an application that runs in WebGL-enabled browsers. After spending some days on it, I finally have something fit to be seen.

This post deals with how I've made that little [random terrain generation application](https://srchea.github.io/Terrain-Generation/) using the [Three.js framework](https://github.com/mrdoob/three.js/) and the [diamond-square algorithm](https://en.wikipedia.org/wiki/Diamond-square_algorithm). Obviously, this works with the new generation of web-browsers (with WebGL support). Latest Google Chrome versions, Mozilla Firefox versions and Internet Explorer 11 are supporting the WebGL context. As always, you can find [the source code](https://github.com/srchea/Terrain-Generation) on [my GitHub profile](https://github.com/srchea/).

Of course, the application gives us more features than applying a texture or showing a mesh. I invite you to discover and to play with it. By the way, I would really appreciate your feedback. 🙂

<p style="text-align: center;">
  <img src="/static/assets/posts/terrain-generation-diamond-square.png" />
</p>

## The diamond-square algorithm

As I said above, I have used the diamond-square algorithm to generate a random fractal terrain. If you already know how the midpoint displacement algorithm works (what if you don't too), the diamond-square algorithm would be pretty easy to understand. Here is a good explanation of these algorithms.

First of all, that algorithm only works on 2D arrays of 2n+1 dimensions (e.g. 129×129, 1025×1025, etc.). As its name suggests, it works on squares and it needs the four corner points, and the midpoint to generate height values regarding average values of corners. After that, it takes the middle of each edges (it actually builds diamonds) and it takes the midpoint to get other squares (sub-squares). [This excellent post by Paul Boxley](http://paulboxley.com/blog/2011/03/terrain-generation-mark-one) shows the behavior of the algorithm step by step.

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

For the texture, Three.js gives the MeshBasicMaterial object to set an image on the mesh with the 'map' attribute. Otherwise, we can put the 'wireframe' and its 'color'.

```javascript
if(this.texture !== null) {
  this.material = new THREE.MeshBasicMaterial({
    map: THREE.ImageUtils.loadTexture(this.texture)
  });
}
else {
  this.material = new THREE.MeshBasicMaterial({
    color : 0x000000,
    wireframe : true
  });
}
```

Subsequently, we build the mesh with its shape and its texture. Then we add it in the scene.

```javascript
this.mesh = new THREE.Mesh(this.geometry, this.material);
this.scene.add(this.mesh);
```

For aesthetic and performance reasons, we can add an fog effect in the whole scene. In fact, the CPU/GPU doesn't have to compute all the terrain to the end but the viewable part. This can be added with:

```javascript
this.scene.fog = new THREE.FogExp2(0xffffff, this.fog);
```

The way of fading fog with distance is controlled by the Exponential Squared mode.

### The camera

The camera that I have used is the perspective projection camera. It is defined by its field of view (fov), aspect ratio, near plane and far plane. In few words, this is similar to the human view.

```javascript
this.fov = 50;
this.aspect = window.innerWidth/window.innerHeight;
this.near = 1;
this.far = 100000;
this.camera = new THREE.PerspectiveCamera(
  this.fov,
  this.aspect,
  this.near,
  this.far
);
```

If you used to play at Call Of Duty, Battlefield, Unreal Tournament, Counter-Strike or you are a FPS gamer, you will find your feet. In fact, the camera is controlled exactly like in a FPS game with just one difference: you need to drag the screen to change the angle. The first person controls offered by the framework doesn't handle correctly what I would like to have, especially for the mouse dragging. However, that lets us to use the keyboard for navigation using the native behavior of the FirstPersonControls (native controls from Three.js), that is to say arrow keys or a (strafe left), w (move up), d (strafe right), s (move down) and r (go upper), f (go lower). So, I have made my own controls, called: `FirstPersonNavigationControls`.

First of all, there are some attributes of the FirstPersonControls that need to be set to allow the FirstPersonNavigationControls to have the wished behavior such as vertical constrains.

```javascript
this.firstPersonControls = new THREE.FirstPersonControls(
  this.object,
  this.domElement
);
this.firstPersonControls.movementSpeed = 5.0;
this.firstPersonControls.lookSpeed = 0.005;
this.firstPersonControls.noFly = true;
this.firstPersonControls.activeLook = false;
this.firstPersonControls.constrainVertical = true;
this.firstPersonControls.verticalMin = 0;
this.firstPersonControls.verticalMax = 0;
```

For the mouse dragging, I have used the longitude/latitude system. It rotates the camera regarding the z axis from 0° to 360° (modulo) and the y axis at 180° (between -90° and 90°).

```javascript
this._lat = Math.max(-90, Math.min(90, this._lat));
this._phi = (90-this._lat)*Math.PI/180;
this._theta = this._lon * Math.PI/180;
this.object.target.x = Math.sin(this._phi)*Math.cos(this._theta);
this.object.target.y = Math.cos(this._phi);
this.object.target.z = Math.sin(this._phi)*Math.sin(this._theta);
```

After normalizing x, y, z almost like in the `lookAt()` method, we obtain the rotation matrix that we apply on the camera. The whole code of the `FirstPersonNavigationControls` class can be found here.

### The rendering

The scene is rendered in the WebGL context and frames updating are fully handled by the framework with the `requestAnimationFrame()` and `render()` functions.

```javascript
this.renderer = new THREE.WebGLRenderer({ antialias: false });
function animate() {
  requestAnimationFrame(animate);
  scene.render();
}
```

## Changeable parameters

The live demo gives a control panel, here is the list of parameters. Variables that directly affect the terrain:

 - __Size__ (width, height)<br />The size of the terrain in pixel.
 - __Number of segments__ (2n+1 dimensions)<br />Note that the more segments you will set, the more computing time it will take.
 - __Smoothing factor__<br />This affects random height variations. The more is the value, the more is the height variation.

Others that affect the scene, without regenerating the terrain:

 - __Textures__<br />You can choose between 3 distinct textures or none: it puts black wireframes.
 - __Fog__<br />The density of the fog used by the exponential squared function (Exp2).
 - __Border__<br />This enables the border on the edges of the terrain.

Debug/performance:

 - __Info__<br />If checked, this displays the camera position/angle and the number of frames per second (FPS) that the browser computes.

## Performance

On the one hand, I have run the application on my laptop under Windows 7, Intel core 2 duo 2.1 GHz and 4 Gb of RAM with Firefox 10 and Chrome 16. Both work with 60 FPS, which is totally correct. On the other hand, I have executed it on my AMD64 1.3 GHz, 3 Gb of RAM using Ubuntu 10.10 and Chrome 17, I got 56 FPS. However, on the same computer, I have some issues to run with Firefox 10: it doesn't update frames.

## Bottom line

That application is not stable, it is still in development. I would really appreciate if you guys send me your performance/bug issues. And feel free to send me some screenshots of your amazing terrain generation. :-)

 - [Random terrain generation (live demo)](https://srchea.github.io/Terrain-Generation/)
 - [Source code on GitHub](https://github.com/srchea/Terrain-Generation)
 - [Three.js framework](https://github.com/mrdoob/three.js/)
