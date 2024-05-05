package data;

public class AdminUser {

  private final String userName;
  private final String password;
  private boolean authenticated;

  public AdminUser(String userName, String password) {
    this.userName = userName;
    this.password = password;
  }

  public String getUserName() {
    return userName;
  }

  public String getPassword() {
    return password;
  }

  public boolean isAuthenticated() {
    return authenticated;
  }

  public void setAuthenticated(boolean authenticated) {
    this.authenticated = authenticated;
  }
}
