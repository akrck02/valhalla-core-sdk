package org.akrck02.valhalla.core.sdk.model.team

import org.akrck02.valhalla.core.sdk.model.user.User
import java.time.LocalDateTime

/**
 * This class represents a team in the app
 */
@Suppress("unused")
data class Team(
    var id: String?,
    var name: String?,
    var owner: User?,
    var description: String?,
    var profilePicturePath: String?,

    var projects: MutableList<Int> = mutableListOf(),
    var members: MutableList<User> = mutableListOf(),

    var creationTime: LocalDateTime?,
    var updateTime: LocalDateTime?
)