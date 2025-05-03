package io.github.solid.resourcepack.bbconv.api.animation.keyframe

import io.github.solid.resourcepack.bbconv.config.PositionKeyframe
import io.github.solid.resourcepack.bbconv.config.RotationKeyframe
import io.github.solid.resourcepack.bbconv.config.ScaleKeyframe
import io.github.solid.resourcepack.bbconv.util.Interpolation
import io.github.solid.resourcepack.bbconv.util.Interpolator
import org.joml.Quaternionf
import org.joml.Vector3f

typealias TimelineFrames<T> = Map<Float, Keyframe<T>>

class Timeline<T>(private val frames: TimelineFrames<T>, private val clazz: Class<T>) {

    private val sortedFrames = frames.toSortedMap(compareBy { it })

    fun interpolate(time: Float): T {
        require(sortedFrames.size >= 2) { "Need at least 2 keyframes, has ${sortedFrames.size}" }
        val times = sortedFrames.keys.toList()

        if (time <= times.first()) return sortedFrames[times.first()]!!.data
        if (time >= times.last()) return sortedFrames[times.last()]!!.data

        val index = times.indexOfLast { it <= time }
        val t0 = times.getOrNull(index - 1)
        val t1 = times[index]
        val t2 = times[index + 1]
        val t3 = times.getOrNull(index + 2)

        val k0 = t0?.let { frames[it] } ?: frames[t1]!!
        val k1 = frames[t1]!!
        val k2 = frames[t2]!!
        val k3 = t3?.let { frames[it] } ?: frames[t2]!!

        val localT = ((time - t1) / (t2 - t1)).coerceIn(0f, 1f)

        val interpolator = Interpolator(k1.interpolation, clazz)
        return interpolator.interpolate(localT, k0.data, k1.data, k2.data, k3.data)
    }

    companion object {
        @JvmStatic
        fun ofPosition(frames: List<PositionKeyframe>): Timeline<Vector3f> {
            val converted = frames.associate {
                it.time to Keyframe(
                    it.position.toVectorf(),
                    Interpolation.fromString(it.interpolation)
                )
            }
            return Timeline(converted, Vector3f::class.java)
        }

        @JvmStatic
        fun ofScale(frames: List<ScaleKeyframe>): Timeline<Vector3f> {
            val converted = frames.associate {
                it.time to Keyframe(
                    it.scale.toVectorf(),
                    Interpolation.fromString(it.interpolation)
                )
            }
            return Timeline(converted, Vector3f::class.java)
        }

        @JvmStatic
        fun ofRotation(frames: List<RotationKeyframe>): Timeline<Quaternionf> {
            val converted = frames.associate {
                it.time to Keyframe(
                    it.leftRotation.toQuaternionf(),
                    Interpolation.fromString(it.interpolation)
                )
            }
            return Timeline(converted, Quaternionf::class.java)
        }
    }
}