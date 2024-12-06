package org.akrck02.valhalla.core.sdk.model.user

import java.time.LocalDateTime

/**
 * This class represents a user in the app
 */
data class User(
    var id: String?,
    var username: String?,
    var email: String?,
    var password: String?,

    var validated: Boolean?,
    var validationCode: String?,

    var profilePicturePath: String?,
    var creationTime: LocalDateTime?
)
