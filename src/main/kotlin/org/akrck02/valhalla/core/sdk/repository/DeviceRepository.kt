package org.akrck02.valhalla.core.sdk.repository

import org.akrck02.valhalla.core.sdk.model.device.Device

/**
 * This interface represents the data access layer
 * for the device collection.
 */
@Suppress("unused")
interface DeviceRepository {

    /**
     * Register a new device for a user
     * @param userId The user id to add the                                                                                                                                        device to
     * @param device The device to register
     * @return The id of the registered device
     */
    suspend fun register(userId: String?, device: Device?): String

    /**
     * Get a device by user and id
     * @param userId The user id
     * @param id The device id
     * @return The requested device
     */
    suspend fun get(userId: String?, id: String?): Device

    /**
     * Get a device by its user and current auth token
     * @param userId The user id
     * @param token The auth token
     * @return The requested device
     */
    suspend fun getByAuth(userId: String?, token: String?): Device

    /**
     * Get all devices by user
     * @param userId The user id
     * @return The devices belonging to a user
     */
    suspend fun getAll(userId: String?): Device

}