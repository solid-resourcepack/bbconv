package io.github.solid.resourcepack.bbconv.api.animation.entity

import io.github.solid.resourcepack.bbconv.api.animation.bone.BoneAnimation
import io.github.solid.resourcepack.bbconv.api.animation.bone.BoneAnimationContext
import io.github.solid.resourcepack.bbconv.api.entity.RenderedBone
import io.github.solid.resourcepack.bbconv.api.entity.RenderedEntity
import io.github.solid.resourcepack.bbconv.config.OpenModelConfig
import io.github.solid.resourcepack.bbconv.util.QuaternionMath
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

    fun get(animation: String): EntityAnimationData? {
        return loadedAnimations.toList().firstOrNull { it.first.name == animation }?.first
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
                            Vector3f(), Quaternionf(), Vector3f(), Quaternionf()
                        )
                    )
                )
            }
            activeAnimations[type] = time
        }
        result.forEach { (bone, transformation) ->
            val transformation = appendTransformation(transformation, bone.getInitialTransformation())
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
        bones: List<RenderedBone>, animation: EntityAnimation
    ): MutableMap<RenderedBone, BoneAnimation> {
        val result = bones.filter { animation.keys.contains(it.bone.id) }
            .associateWith { bone -> animation.toList().first { it.first == bone.bone.id }.second }.toMutableMap()
        bones.forEach { bone ->
            result.putAll(bonesToAnimations(bone.children, animation))
        }
        return result
    }

    private fun partialAnimate(
        animation: BoneAnimation, ctx: BoneAnimationContext, transformation: Transformation
    ): Map<RenderedBone, Transformation> {
        val result = appendTransformation(transformation, animation.animate(ctx))
        val returned = mutableMapOf<RenderedBone, Transformation>()
        returned[ctx.self] = result
        returned.putAll(animateChildren(result, ctx.self))
        return returned
    }

    private fun animateChildren(
        transformation: Transformation, parent: RenderedBone
    ): Map<RenderedBone, Transformation> {
        val result = mutableMapOf<RenderedBone, Transformation>()
        parent.children.forEach { child ->
            val initial = child.getInitialTransformation()
            val parentTrans = transformation

            // Calculate rotated local translation
            val rotatedTranslation = Vector3f(initial.translation)
            parentTrans.leftRotation.transform(rotatedTranslation)

            val deltaTranslation = Vector3f(rotatedTranslation).sub(initial.translation)

            // Rotation delta: how the rotation has changed
            val rotatedRotation = Quaternionf(parentTrans.leftRotation).mul(initial.leftRotation)
            val deltaRotation = QuaternionMath.delta(initial.leftRotation, rotatedRotation)

            // Scale delta: how the scale has changed
            val deltaScale = Vector3f(parentTrans.scale).div(parent.bone.scale).mul(initial.scale)

            val childTransformation = Transformation(
                deltaTranslation,
                deltaRotation,
                deltaScale,
                Quaternionf()
            )

            result[child] = childTransformation
            result.putAll(animateChildren(childTransformation, child))
        }
        return result
    }

    private fun appendTransformation(first: Transformation, second: Transformation): Transformation {
        return Transformation(
            Vector3f(first.translation).add(second.translation),
            Quaternionf(first.leftRotation).mul(Quaternionf(second.leftRotation)),
            Vector3f(first.scale).add(Vector3f(second.scale)),
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