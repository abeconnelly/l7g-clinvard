var x = lvcvd_sql("select id, clinvar_id, tileID from clinvar_tilemap limit 10;");
var x_json = JSON.parse(x);
vard_return(x_json, "  ");
