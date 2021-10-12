package Question2;

import javafx.application.Application;
import javafx.event.ActionEvent;
import javafx.event.EventHandler;
import javafx.geometry.Insets;
import javafx.scene.Scene;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.TextField;
import javafx.scene.layout.GridPane;
import javafx.scene.layout.BorderPane;
import javafx.scene.layout.HBox;
import javafx.stage.Stage;
import javafx.scene.control.TextArea;

public class Question2 extends Application {
    
    public String text;
    
   public void start(Stage primaryStage){
   primaryStage.setTitle("Question-2");
   primaryStage.show();
   TextArea txt1=new TextArea();
   TextArea txt2=new TextArea();
   double height = 400; 
   double width = 300;  
txt1.setPrefHeight(height);
txt1.setPrefWidth(width); 
   
   
   HBox hbox = new HBox(txt1,txt2);
   
   Scene scene = new Scene(hbox,800,600);
   primaryStage.setScene(scene);
   primaryStage.show();
   
   }
    public static void main(String[] args) {
        launch(args);
    }
}
