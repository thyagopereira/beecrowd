
function search(nums: number[], target: number): number {
    var l:number = 0, r:number = nums.length -1;

    while(l <= r){
        var mid:number = Math.round( (l + r) / 2);

        if(target === nums[mid]){
            return mid;
        }
        // Left Sorted Portion
        if(nums[l] <= nums[mid]){

            if((target >= nums[mid]) || (target < nums[l])){
                l = mid + 1;
            }else{
                r = mid - 1;
            }

        // Rigth Sorted Portion
        }else{

            if ((target < nums[mid]) || (target > nums[r])){
                r = mid - 1;
            }else{
                l = mid + 1;
            }
        }
    }

    return -1;
};
