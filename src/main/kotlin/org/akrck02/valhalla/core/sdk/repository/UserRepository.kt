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
     */
    fun register(user: User ?)

    /**
     * Get the user by id
     * @param id The user id
     * @return The requested user
     */
    fun get(id : String ?) : User ?

    /**
     * Get the user by email
     * @param email The user email
     * @return The requested user
     */
    fun getByEmail(email : String ?) : User ?

    /**
     * Update the user data
     * @param id The user id
     * @param user The new user data
     */
    fun update(id : String ?, user: User ?)

    /**
     * Update the profile picture for a user
     * @param id The user id
     * @param picture The profile picture to be set
     */
    fun updateProfilePicture(id : String ?, picture : ByteArray ?)

    /**
     * Delete the user
     * @param id
     */
    fun delete(id : String ?)

    /**
     * Login with a user in the service
     * @param user The user that wants to log in
     * @param device The device that the user is using
     * @return The auth token
     */
    fun login(user: User ?, device : Device ?) : String?

    /**
     * Login the platform with the auth token representing
     * a user with a device
     * @param id The id of user that wants to log in
     * @param token The auth token used
     */
    fun loginWithAuth(id: String ?, token : String ?)

    /**
     * Validate the user account
     * @param code The validation code
     */
    fun validateAccount(code : String?)
}