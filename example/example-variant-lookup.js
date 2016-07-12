var query = "select id, chrom, pos, rsid, ref, alt, qual, filter, info from clinvar where info like '%RCV000001485.5%'  limit 10 ;"
var r_json = JSON.parse(lvcvd_sql(query));
vard_return(r_json, "  ");
