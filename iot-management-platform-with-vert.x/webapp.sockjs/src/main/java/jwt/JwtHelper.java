package jwt;

import io.vertx.core.json.JsonObject;
import io.vertx.ext.auth.JWTOptions;
import io.vertx.ext.auth.PubSecKeyOptions;
import io.vertx.ext.auth.jwt.JWTAuthOptions;
import io.vertx.rxjava3.core.Vertx;
import io.vertx.rxjava3.ext.auth.jwt.JWTAuth;
import io.vertx.rxjava3.ext.web.handler.JWTAuthHandler;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

public class JwtHelper {

  private static JwtHelper INSTANCE;
  private final JWTAuth provider;
  private final JWTAuthHandler handler;

  private JwtHelper(Vertx vertx, List<String> keyFiles) {
    List<String> permFiles = keyFiles.stream().map(file -> {
      try {
        return Files.readString(Paths.get(file));
      } catch (IOException e) {
        throw new RuntimeException(e);
      }
    }).toList();

    this.provider = JWTAuth.create(vertx, new JWTAuthOptions()
      .addPubSecKey(new PubSecKeyOptions()
        .setAlgorithm("RS256")
        .setBuffer(permFiles.get(0)))
      .addPubSecKey(new PubSecKeyOptions()
        .setAlgorithm("RS256")
        .setBuffer(permFiles.get(1)))
    );
    this.handler = JWTAuthHandler.create(this.provider);
  }

  public static JwtHelper getInstance(Vertx vertx, List<String> keyFiles) {
    if (INSTANCE == null) {
      INSTANCE = new JwtHelper(vertx, keyFiles);
    }
    return INSTANCE;
  }

  public JWTAuth getProvider() {
    return provider;
  }

  public JWTAuthHandler getHandler() {
    return handler;
  }

  public Object generateToken(String username, JsonObject dataPayload) {
    var jwtOptions = new JWTOptions()
      .setAlgorithm("RS256")
      .setExpiresInMinutes(600)
      .setIssuer("https://github.com/okmich")
      .setSubject(username);
    return this.provider.generateToken(dataPayload, jwtOptions);
  }

}
