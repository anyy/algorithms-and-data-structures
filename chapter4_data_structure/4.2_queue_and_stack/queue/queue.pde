import java.util.*;

Queue qu = new ArrayDeque();
qu.add(1);
qu.add(2);
qu.add(3);
qu.add(4);

// output using iterator
Iterator it = qu.iterator();
while (it.hasNext()) {
  println(it.next());
}

// output using poll
println(qu.poll());
println(qu.poll());
println(qu.poll());
