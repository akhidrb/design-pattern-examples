package patterns.decorator;

public abstract class Decorator implements VisualComponent {

    protected VisualComponent decoratedComponent;

    public Decorator(VisualComponent decoratedComponent) {
        super();
        this.decoratedComponent = decoratedComponent;
    }

}
