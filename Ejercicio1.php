+<?php
+
+
+for($num=1;$num<=100;$num++){
+
+        if ($num%3==0)  {
+                echo "CD";
+        }
+        if ($num%5==0){
+                echo "mon";
+        }
+        if ($num%3!=0 && $num%5!=0){
+                echo $num;
+        }
+        echo ", ";
+}
+
+?>
