package org.akrck02.valhalla.core.sdk.repository

import org.akrck02.valhalla.core.sdk.model.device.Device

/**
 * This interface represents the data access layer
 * for the device collection.
 */
@Suppress("unused")
interface DeviceRepository {

    /**
     * Register a new device
     * @param device The device to register
     * @return The id of the registered device
     */
    suspend fun register(device: Device?): String

    /**
     * Get a device by id
     * @param id The device id
     * @return The requested device
     */
    suspend fun get(id: String?): Device

    /**
     * Get a device by its current auth token
     * @param token The auth token
     * @return the requested device
     */
    suspend fun getByAuth(token: String?): Device

    /**
     * Get all the devices belonging to a user
     * @param userId The owner user id
     * @return The list of devices
     */
    suspend fun getAllByOwner(userId: String?): List<Device>

    /**
     * Update a device
     * @param id The device id
     * @param device The new device data
     */
    suspend fun update(id: String?, device: Device?)

}