import org.openqa.selenium.*;
import org.openqa.selenium.chrome.ChromeDriver;


public class RqThruBrowser {

    public static final String DRIVER_LINK = BaseProperties.getInstance().getProperties().getProperty("webdriver.chrome.driver");

    public static final String API_URL = BaseProperties.getInstance().getProperties().getProperty("api.url");
    public static final String TOKEN = BaseProperties.getInstance().getProperties().getProperty("bot.token");
    public static final String GET_UPDATES = API_URL + TOKEN + "/getUpdates";
    public static final String SEND_MESSAGE = API_URL + TOKEN + "/sendMessage";

    private static WebDriver driver;

    public void setup() {
        System.setProperty("webdriver.chrome.driver", DRIVER_LINK);
        driver = new ChromeDriver();
        driver.manage().window().maximize();
        //driver.manage().timeouts().implicitlyWait(3, TimeUnit.SECONDS);
    }

    public String updateMessages() {
        driver.get(GET_UPDATES);
        return getWebElementFromDriver("xpath", "//pre").getText();
    }

    public String sendMessage(String messageText, String chatID) {
        driver.get(SEND_MESSAGE + "?chat_id=" + chatID + "&text=" + messageText);
        return messageText;
    }

    private WebElement getWebElementFromDriver(String type, String value) {
        if (type.equals("id")) {
            return driver.findElement(By.id(value));
        } else if (type.equals("xpath")) {
            return driver.findElement(By.xpath(value));
        } else if (type.equals("className")) {
            return driver.findElement(By.className(value));
        } else if (type.equals("tagName")) {
            return driver.findElement(By.tagName(value));
        } else if (type.equals("linkText")) {
            return driver.findElement(By.linkText(value));
        } else {
            return null;
        }
    }

    public void tearDown() {
        driver.quit();
    }
}