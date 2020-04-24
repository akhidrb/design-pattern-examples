package patterns.decorator.decorators;

import patterns.decorator.Decorator;
import patterns.decorator.VisualComponent;

public class ScrollDecorator extends Decorator {

    public ScrollDecorator(VisualComponent visualComponent) {
        super(visualComponent);
    }

    @Override
    public void draw() {
        System.out.println("Top Scroll");
        decoratedComponent.draw();
        System.out.println("Bottom Scroll");
    }

}
