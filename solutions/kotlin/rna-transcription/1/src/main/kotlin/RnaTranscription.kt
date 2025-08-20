fun transcribeToRna(dna: String): String {
    var rna=StringBuilder()
    for(nucleotide in dna)
    {
        when(nucleotide)
        {
            'A' ->rna.append('U')
            'G'-> rna.append('C')
            'T'-> rna.append('A')
            else-> rna.append('G')
        }
    }
    return rna.toString()
   
}
