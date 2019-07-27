import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class RqBotStubServlet extends HttpServlet {

    public static final String GETUPDATES = "getupdates";
    public static final String SENDMESSAGE = "sendmessage";

    private final RqThruBrowser rqThruBrowser;

    public RqBotStubServlet(RqThruBrowser rqThruBrowser) {
        this.rqThruBrowser = rqThruBrowser;
    }

    public void doGet(HttpServletRequest request,
                      HttpServletResponse response) throws ServletException, IOException {
        String action = request.getParameter("action");
        String rsDirectlyFromBot = "";
        if (GETUPDATES.equals(action)) {
            rsDirectlyFromBot = rqThruBrowser.updateMessages();
        } else if (SENDMESSAGE.equals(action)) {
            String message = request.getParameter("value");
            String chatID = request.getParameter("chatid");
            rsDirectlyFromBot = rqThruBrowser.sendMessage(message, chatID);
        }
        response.getWriter().println(rsDirectlyFromBot);
        response.setContentType("text/html;charset=utf-8");
        response.setStatus(HttpServletResponse.SC_OK);
    }
}
