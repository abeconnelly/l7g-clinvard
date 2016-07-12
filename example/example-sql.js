var x = lvcvd_sql("select * from clinvar limit 10;");
var x_json = JSON.parse(x);
vard_return(x_json, "  ");
