package com.example.alertbridge.infrastructure.redis;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AlertStateRedisCrudRepository extends CrudRepository<AlertStateRedis, String> {
}
