package org.akrck02.valhalla.core.sdk.repository

import org.akrck02.valhalla.core.sdk.model.device.Device
import org.akrck02.valhalla.core.sdk.model.user.User

/**
 * This interface represents the data access layer
 * for the user collection.
 */
@Suppress("unused")
interface UserRepository {

    /**
     * Register a user
     * @param user The user to register
     * @return The id of the registered user
     */
    suspend fun register(user: User?): String

    /**
     * Get the user by id
     * @param id The user id
     * @param secure If the sensible data must be hidden
     * @return The requested user
     */
    suspend fun get(id: String?, secure: Boolean? = true): User

    /**
     * Get the user by email
     * @param email The user email
     * @param secure If the sensible data must be hidden
     * @return The requested user
     */
    suspend fun getByEmail(email: String?, secure: Boolean? = true): User

    /**
     * Update the user data
     * @param id The user id
     * @param user The new user data
     */
    suspend fun update(id: String?, user: User?)

    /**
     * Update the profile picture for a user
     * @param id The user id
     * @param picture The profile picture to be set
     */
    suspend fun updateProfilePicture(id: String?, picture: ByteArray?)

    /**
     * Delete the user
     * @param id
     */
    suspend fun delete(id: String?)

    /**
     * Login with a user in the service
     * @param user The user that wants to log in
     * @param device The device that the user is using
     * @return The auth token
     */
    suspend fun login(user: User?, device: Device?): String

    /**
     * Login the platform with the auth token representing
     * a user with a device
     * @param id The id of user that wants to log in
     * @param token The auth token used
     */
    suspend fun loginWithAuth(id: String?, token: String?)

    /**
     * Validate the user account
     * @param code The validation code
     */
    suspend fun validateAccount(code: String?)
}