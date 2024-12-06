package org.akrck02.valhalla.core.sdk.model.project

import org.akrck02.valhalla.core.sdk.model.team.Team
import org.akrck02.valhalla.core.sdk.model.user.User
import java.time.LocalDateTime

/**
 * This class represent a project in the app
 */
@Suppress("unused")
data class Project(
    var id: String?,
    var name: String?,
    var description: String?,
    var owner: User?,
    var teams: MutableList<Team> = mutableListOf(),
    var creationTime: LocalDateTime?,
    var updateTime: LocalDateTime?
)