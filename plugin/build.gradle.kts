import com.github.jengelman.gradle.plugins.shadow.tasks.ShadowJar
import dev.s7a.gradle.minecraft.server.tasks.LaunchMinecraftServerTask

plugins {
    id("dev.s7a.gradle.minecraft.server") version "3.2.1"
}

dependencies {
    implementation(project(":api"))
}

tasks.named("shadowJar", ShadowJar::class) {
    mergeServiceFiles()
    archiveFileName = "${rootProject.name}.jar"
}

task<LaunchMinecraftServerTask>("launchMinecraftServer") {
    dependsOn("shadowJar")

    doFirst {
        copy {
            from(layout.buildDirectory.file("libs/${rootProject.name}.jar"))
            into(layout.buildDirectory.file("MinecraftServer/plugins"))
        }
    }
    jarUrl.set(LaunchMinecraftServerTask.JarUrl.Paper("1.21.5"))
    agreeEula.set(true)
}

kotlin {
    jvmToolchain(21)
}