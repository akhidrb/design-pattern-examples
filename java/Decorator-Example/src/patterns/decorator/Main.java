package patterns.decorator;

import patterns.decorator.components.TextView;
import patterns.decorator.decorators.BorderDecorator;
import patterns.decorator.decorators.ScrollDecorator;

public class Main {

    public static void main(String[] args) {

        VisualComponent visualComponent = new BorderDecorator(new ScrollDecorator(new TextView()));
        visualComponent.draw();
    }
}
