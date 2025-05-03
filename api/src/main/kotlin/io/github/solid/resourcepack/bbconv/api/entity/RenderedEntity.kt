package io.github.solid.resourcepack.bbconv.api.entity

import io.github.solid.resourcepack.bbconv.api.animation.entity.EntityAnimationController
import io.github.solid.resourcepack.bbconv.config.Bone
import io.github.solid.resourcepack.bbconv.config.OpenModelConfig
import org.bukkit.Location
import java.util.*

data class RenderedEntity(
    val location: Location,
    val config: OpenModelConfig,
    val id: UUID = UUID.randomUUID(),
) {
    private val renderedBones = mutableMapOf<String, RenderedBone>()
    val rootBones = mutableListOf<RenderedBone>()

    private val animationController = EntityAnimationController(this)

    init {
        location.pitch = 0f
        config.boneTree.forEach { bone ->
            val bone = initBone(bone) ?: return@forEach
            rootBones.add(bone)
        }
        if (rootBones.isEmpty()) {
            findRootBones(config.boneTree)
        }
    }

    private fun findRootBones(bones: List<Bone>) {
        var found = false
        for (bone in bones) {
            bone.children.forEach {
                if (renderedBones.containsKey(it.id)) {
                    rootBones.add(renderedBones[it.id]!!)
                    found = true
                }
            }
        }
        if (!found) {
            bones.forEach { bone ->
                findRootBones(bone.children)
            }
        }
    }

    private fun initBone(bone: Bone, parent: RenderedBone? = null): RenderedBone? {
        var rendered: RenderedBone? = null
        if (bone.visible) {
            rendered = RenderedBone(this, bone)
            parent?.children?.add(rendered)
            renderedBones[bone.id] = rendered
        }
        bone.children.forEach { child ->
            initBone(child, rendered)
        }
        return rendered
    }

    fun getAnimationController(): EntityAnimationController {
        return animationController
    }

    fun spawn() {
        renderedBones.values.forEach { it.spawn() }
    }


}