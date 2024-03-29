package com.okmich.livemanning.gridmonitoring.data.repo;

import com.okmich.livemanning.gridmonitoring.data.entities.Device;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface DeviceRepository extends JpaRepository<Device, String> {

}
