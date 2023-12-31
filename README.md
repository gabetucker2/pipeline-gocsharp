# pipeline-gocsharp
This is a simple repository intended to store a prototype codebase for another project.

My emerstack project relies on go-to-C# communication so that Unity and the neural network can communicate.  This codebase is the testing grounds which lays the groundwork pipeline for sending signals from C# - CSV - Go and vice versa.

Here is a diagram for the schematic:

![](schematic.png)

The DynamicsGo folder is "Go" in the schematic; likewise for the DynamicsUnity folder.  The Translator folder is "CSV" in the schematic.

If you haven't already, install [Go](https://go.dev/doc/install) and [Unity Hub](https://unity.com/download).  Then, click ```Open``` in Unity Hub and select your DynamicsUnity folder.  After doing so, it should prompt you to install the proper Unity version to access DynamicsUnity:

![](unityVersion.png)

Make sure to run Unity before running Go commands:

![](unityStart.png)

Finally, to control the people through Go, do:

```C:\...\pipeline-gocsharp\DynamicsGo> go run . Bed NA```

This will instruct Franklin to the Bed and Jerry to stay stationary.
