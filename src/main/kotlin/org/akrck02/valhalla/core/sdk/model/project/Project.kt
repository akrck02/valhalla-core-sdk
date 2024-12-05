package org.akrck02.valhalla.core.sdk.model.project

import org.akrck02.valhalla.core.sdk.model.team.Team
import org.akrck02.valhalla.core.sdk.model.user.User
import java.time.LocalDateTime

/**
 * This class represent a project in the app
 */
@Suppress("unused")
data class Project(
    val id: String?,
    var name: String? = null,
    var description: String? = null,
    var owner: User? = null,

    var teams: MutableList<Team> = mutableListOf(),

    var creationTime: LocalDateTime? = null,
    var updateTime: LocalDateTime? = null
)