package patterns.decorator.decorators;

import patterns.decorator.Decorator;
import patterns.decorator.VisualComponent;

public class BorderDecorator extends Decorator {

    public BorderDecorator(VisualComponent visualComponent) {
        super(visualComponent);
    }

    @Override
    public void draw() {
        System.out.println("Top Border");
        decoratedComponent.draw();
        System.out.println("Bottom Border");
    }


}
