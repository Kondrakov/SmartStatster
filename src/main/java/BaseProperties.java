import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.util.Properties;

public class BaseProperties {
    private final Properties properties = new Properties();
    private static BaseProperties INSTANCE = null;

    private BaseProperties(){
        try {
            properties.load(new FileInputStream(new File("src/main/resources/config/"
                    + System.getProperty("environment", "default") + ".properties")));
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static BaseProperties getInstance() {
        if (INSTANCE == null){
            INSTANCE = new BaseProperties();
        }
        return INSTANCE;
    }

    public Properties getProperties() {
            return properties;
        }
}
