import dev.s7a.gradle.minecraft.server.tasks.LaunchMinecraftServerTask

plugins {
    kotlin("jvm") version "2.1.10"
    id("com.gradleup.shadow") version "8.3.3"
    id("dev.s7a.gradle.minecraft.server") version "3.2.1"
}

group = "io.solid-resourcepack.bbconv"
version = "0.1.0"

repositories {
    mavenCentral()
    maven {
        name = "papermc"
        url = uri("https://repo.papermc.io/repository/maven-public/")
    }
}

dependencies {
    testImplementation(kotlin("test"))
    implementation("org.spongepowered:configurate-extra-kotlin:4.1.2")
    implementation("org.spongepowered:configurate-gson:4.1.2")
    compileOnly("io.papermc.paper:paper-api:1.21.4-R0.1-SNAPSHOT")
}

tasks.test {
    useJUnitPlatform()
}

task<LaunchMinecraftServerTask>("launchMinecraftServer") {
    dependsOn("shadowJar")

    doFirst {
        copy {
            from(layout.buildDirectory.file("libs/${project.name}-${project.version}-all.jar"))
            into(layout.buildDirectory.file("MinecraftServer/plugins"))
        }
    }
    jarUrl.set(LaunchMinecraftServerTask.JarUrl.Paper("1.21.4"))
    agreeEula.set(true)
}

kotlin {
    jvmToolchain(21)
}