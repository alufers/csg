import sys
import bpy

argv = sys.argv
argv = argv[argv.index("--") + 1:]

imported_object = bpy.ops.import_scene.obj(filepath=argv[0])

objs = bpy.data.objects
objs.remove(objs["Cube"], True)

bpy.ops.view3d.camera_to_view_selected()


for area in bpy.context.screen.areas:
    if area.type == 'VIEW_3D':
        area.spaces[0].region_3d.view_perspective = 'CAMERA'
        break
scn = bpy.context.scene
scn.render.filepath = argv[1]
     # render opengl
bpy.ops.render.opengl(write_still=True)

#bpy.ops.wm.quit_blender()
