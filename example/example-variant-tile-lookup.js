var result = {};

var clinvar_query = "select id, chrom, pos, rsid, ref, alt, qual, filter, info from clinvar where info like '%RCV000001485.5%'  limit 10 ;"
var ret_cv_json = JSON.parse(lvcvd_sql(clinvar_query));
var clinvar_id = parseInt(ret_cv_json.result[1][0]);



result.clinvar_id = clinvar_id;
result.clinvar_pos = ret_cv_json.result[1][2];
result.rsid = ret_cv_json.result[1][3];
result.clinvar_ref = ret_cv_json.result[1][4];
result.clinvar_alt = ret_cv_json.result[1][5];
result.clinvar_info = ret_cv_json.result[1][8];

var tile_query = "select id, clinvar_id, tileID from clinvar_tilemap where clinvar_id = " + clinvar_id + ";";
var ret_t_json = JSON.parse(lvcvd_sql(tile_query));

result.tile_info = [];

for (var i=1; i<ret_t_json.result.length; i++) {
  var t_inf = {};
  t_inf.tileID = ret_t_json.result[i][2];

  var tile_parts = t_inf.tileID.split(".");
  var tilepath = parseInt(tile_parts[0], 16);
  var tilestep = parseInt(tile_parts[2], 16);
  var tilevar  = parseInt(tile_parts[3], 16);

  var assembly_query = "select reference_name, chromosome, tilepath, tilestep, reference_start, reference_length from lightning_tile_assembly where tilepath = " + tilepath + " and tilestep = " + tilestep + ";";
  var ret_a = JSON.parse(lvcvd_sql(assembly_query));

  t_inf.refName = ret_a.result[1][0];
  t_inf.refChrom = ret_a.result[1][1];
  t_inf.refStart = ret_a.result[1][4];
  t_inf.refLen = ret_a.result[1][5];

  result.tile_info.push(t_inf);

}


vard_return(result, "  ");
