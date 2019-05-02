import org.eclipse.jetty.server.Handler;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.server.handler.HandlerList;
import org.eclipse.jetty.server.handler.ResourceHandler;
import org.eclipse.jetty.servlet.ServletContextHandler;
import org.eclipse.jetty.servlet.ServletHolder;

public class Main {

    public static void main(String[] args) {
        RqThruBrowser rqThruBrowser = new RqThruBrowser();
        rqThruBrowser.setup();
        //rqThruBrowser.tearDown();

        ServletContextHandler context = new ServletContextHandler(ServletContextHandler.SESSIONS);
        context.addServlet(new ServletHolder(new RqBotStubServlet(rqThruBrowser)), "/bot_get");
        ResourceHandler resource_handler = new ResourceHandler();
        HandlerList handlers = new HandlerList();
        handlers.setHandlers(new Handler[]{resource_handler, context});
        Server server = new Server(8085);
        server.setHandler(handlers);
        try {
            server.start();
            server.join();
        } catch (Exception ex) {
            ex.printStackTrace();
        }
    }
}
