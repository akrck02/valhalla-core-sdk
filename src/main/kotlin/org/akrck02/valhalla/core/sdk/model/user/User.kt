package org.akrck02.valhalla.core.sdk.model.user

import java.time.LocalDateTime

/**
 * This class represents a user in the app
 */
data class User(val id: String?) {

    var username: String? = null
    var password: String? = null

    var validated: Boolean? = null
    var validationCode: String? = null

    var profilePicturePath: String? = null
    var creationTime: LocalDateTime? = null

}
