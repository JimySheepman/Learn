package nLayerDemo.core;

import nLayerDemo.jLogger.JLoggerManager;

// başakasının yazdığı kodu adepte ediyoruz.
// kendi interfacesimize yönelik yeniden yazarak diğer servisleri çağrıyoruz
public class JLoggerManagerAdapter implements LoggerService {
    @Override
    public void logToSystem(String message) {
        JLoggerManager manager = new JLoggerManager();
        manager.log(message);
    }
}
