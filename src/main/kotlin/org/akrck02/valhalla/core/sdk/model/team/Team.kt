package org.akrck02.valhalla.core.sdk.model.team

import org.akrck02.valhalla.core.sdk.model.user.User
import java.time.LocalDateTime

/**
 * This class represents a team in the app
 */
@Suppress("unused")
data class Team(
    var id: String? = null,
    var name: String? = null,
    var owner: User? = null,
    var description: String? = null,
    var profilePicturePath: String? = null,

    var projects: MutableList<Int> = mutableListOf(),
    var members: MutableList<User> = mutableListOf(),

    var creationTime: LocalDateTime? = null,
    var updateTime: LocalDateTime? = null,
)