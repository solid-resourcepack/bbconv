package io.github.solid.resourcepack.bbconv.api.animation.entity

import io.github.solid.resourcepack.bbconv.api.animation.bone.BoneAnimation
import io.github.solid.resourcepack.bbconv.api.animation.bone.BoneAnimationContext
import io.github.solid.resourcepack.bbconv.api.entity.RenderedBone
import io.github.solid.resourcepack.bbconv.api.entity.RenderedEntity
import io.github.solid.resourcepack.bbconv.config.OpenModelConfig
import org.bukkit.util.Transformation
import org.joml.Quaternionf
import org.joml.Vector3f

class EntityAnimationController(
    private val entity: RenderedEntity,
) {
    val loadedAnimations = AnimationInfos.of(entity.config).toMutableMap()
    private val activeAnimations: MutableMap<EntityAnimationData, Float> = mutableMapOf()

    fun play(animation: EntityAnimation) {
        val type = loadedAnimations.toList().firstOrNull { it.second == animation }?.first ?: return
        activeAnimations[type] = 0f
    }

    fun play(animation: String) {
        val type = loadedAnimations.toList().firstOrNull { it.first.name == animation }?.first ?: return
        activeAnimations[type] = 0f
    }

    fun cancel(animation: EntityAnimation) {
        val type = loadedAnimations.toList().firstOrNull { it.second == animation }?.first ?: return
        activeAnimations.remove(type)
    }

    fun cancel(animation: String) {
        val type = loadedAnimations.toList().firstOrNull { it.first.name == animation }?.first ?: return
        activeAnimations.remove(type)
    }

    fun reset() {
        activeAnimations.clear()
    }

    fun animate(delta: Float) {
        val result = mutableMapOf<RenderedBone, Transformation>()
        activeAnimations.forEach { (type, currentTime) ->
            val animation = loadedAnimations[type] ?: return@forEach
            var time = currentTime + delta
            if (time > type.duration) {
                if (!type.looped) {
                    activeAnimations.remove(type)
                    return@forEach
                }
                time = 0f;
            }
            val mappedBones = bonesToAnimations(entity.rootBones, animation)
            mappedBones.forEach { (bone, anim) ->
                result.putAll(
                    partialAnimate(
                        anim, createContext(bone, time), result[bone] ?: Transformation(
                            Vector3f(), Quaternionf(), Vector3f(),
                            Quaternionf()
                        )
                    )
                )
            }
            activeAnimations[type] = time
        }
        result.forEach { (bone, transformation) ->
            val transformation = appendTransformation(bone.initialTransformation, transformation)
            println("changes in bone ${bone.bone.id}: $transformation")
            bone.getDisplay().transformation = transformation
        }
    }

    private fun createContext(bone: RenderedBone, time: Float): BoneAnimationContext {
        return BoneAnimationContext(
            entity,
            time,
            bone,
        )
    }

    private fun bonesToAnimations(
        bones: List<RenderedBone>,
        animation: EntityAnimation
    ): MutableMap<RenderedBone, BoneAnimation> {
        val result = bones.filter { animation.keys.contains(it.bone.id) }
            .associateWith { bone -> animation.toList().first { it.first == bone.bone.id }.second }
            .toMutableMap()
        bones.forEach { bone ->
            result.putAll(bonesToAnimations(bone.children, animation))
        }
        return result
    }

    private fun partialAnimate(
        animation: BoneAnimation,
        ctx: BoneAnimationContext,
        transformation: Transformation
    ): Map<RenderedBone, Transformation> {
        val result = appendTransformation(transformation, animation.animate(ctx))
        val returned = mutableMapOf<RenderedBone, Transformation>()
        returned[ctx.self] = result
        returned.putAll(animateChildren(result, ctx.self))
        return returned
    }

    private fun animateChildren(
        transformation: Transformation,
        parent: RenderedBone
    ): Map<RenderedBone, Transformation> {
        val result = mutableMapOf<RenderedBone, Transformation>()
        parent.children.forEach { child ->

            val initial = child.initialTransformation
            val parentTrans = transformation

            val localTranslation = Vector3f(initial.translation)
                .mul(parentTrans.scale)
                .rotate(parentTrans.leftRotation)

            val worldTranslation = Vector3f(parentTrans.translation).add(localTranslation)

            val worldRotation = Quaternionf(parentTrans.leftRotation).mul(initial.leftRotation)

            val worldScale = Vector3f(parentTrans.scale).mul(initial.scale)

            val childTransformation = Transformation(
                worldTranslation,
                worldRotation,
                worldScale,
                Quaternionf()
            )
            result[child] = childTransformation
            result.putAll(animateChildren(childTransformation, child))
        }
        return result
    }

    private fun appendTransformation(first: Transformation, second: Transformation): Transformation {
        return Transformation(
            first.translation.add(second.translation),
            Quaternionf(second.leftRotation).mul(first.leftRotation),
            first.scale.add(second.scale),
            Quaternionf()
        )
    }
}

object AnimationInfos {

    private val cache = mutableMapOf<OpenModelConfig, Map<EntityAnimationData, EntityAnimation>>()

    @JvmStatic
    fun of(config: OpenModelConfig): Map<EntityAnimationData, EntityAnimation> {
        if (cache.containsKey(config)) {
            return cache[config]!!
        }
        val result = config.animations.associate { animation ->
            EntityAnimationData.of(animation) to EntityAnimations.of(animation)
        }
        cache[config] = result
        return result
    }

    fun clearCache() {
        cache.clear()
    }
}