package patterns.decorator.components;

import patterns.decorator.VisualComponent;

public class TextView implements VisualComponent {

    @Override
    public void draw() {
        System.out.println("Draw Text");
    }

}
