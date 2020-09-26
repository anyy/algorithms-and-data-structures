import java.util.*;

Stack st = new Stack();

// push
st.push(1);
st.push(2);
st.push(3);
st.push(4);

// output
for (int i=0; i<st.size(); i++) {
  println(st.get(i));
}

// output using pop
println(st.pop());

// output using iterator
Iterator it = st.iterator();
while (it.hasNext()) {
  println(it.next());
}
