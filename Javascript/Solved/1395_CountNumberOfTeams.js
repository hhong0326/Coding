/**
 * @param {number[]} rating
 * @return {number}
 */
 var numTeams = function(rating) {
    
    let count = 0;
    for (let i=0; i<rating.length; i++) {
        let t1 = rating[i];
        
        for  (let j=i+1; j<rating.length; j++) {
            if (t1 < rating[j]) {
                let t2 = rating[j];
                for  (let k=j+1; k<rating.length; k++) {
                    if (t2 < rating[k]) {
                        count++;
                    }
                }
            } else if (t1 > rating[j]){
                let t2 = rating[j];
                for  (let k=j+1; k<rating.length; k++) {
                    if (t2 > rating[k]) {
                        count++;
                    }
                }
            }
            
        }
    }
    
    return count;
};