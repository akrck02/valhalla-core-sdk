package org.valhalla.core.sdk.repository

import org.valhalla.core.sdk.model.device.Device
import org.valhalla.core.sdk.model.exception.ServiceException
import org.valhalla.core.sdk.model.user.User

/** This interface represents the data access layer for the users.*/
@Suppress("unused")
interface UserRepository {

    /**
     * Register a new [User] returning the id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun register(user: User?): String

    /**
     * Get the [User] by id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun get(id: String?, secure: Boolean? = true): User

    /**
     * Get the [User] by email.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun getByEmail(email: String?, secure: Boolean? = true): User

    /**
     * Update the [User] data given id the user id and the new data.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun update(id: String?, user: User?)

    /**
     * Update the profile picture for a [User] given the user id and picture to be set.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun updateProfilePicture(id: String?, picture: ByteArray?)

    /**
     * Delete a [User] given the id.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun delete(id: String?)

    /**
     * Login with a [User] in the platform returning the auth token.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun login(user: User?, device: Device?): String

    /**
     * Login into the platform with the auth token representing a [User] with a device.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun loginWithAuth(id: String?, token: String?)

    /**
     * Validate the [User] account given the validation code.
     * Throws [ServiceException] if an error occurs.
     */
    suspend fun validateAccount(code: String?)
}