var x = lvcvd_sql("select id, reference_name, chromosome, tilepath, tilestep, reference_start, reference_length  from lightning_tile_assembly limit 10;");
var x_json = JSON.parse(x);
vard_return(x_json, "  ");
