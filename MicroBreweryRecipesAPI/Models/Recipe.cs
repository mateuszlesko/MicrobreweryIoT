using System.Collections.Generic;
namespace MicroBreweryRecipesAPI.Models
{
    public class Recipe
    {
        //entity proprties
        public int Id {get;set;}
        public string RecipeName {get;set;}


        //relationships
        public int RecipeCategoryId {get;set;}
        public RecipeCategory RecipeCategory {get;set;}
        public List<RecipeStage> RecipeStages {get;set;}
        public List<RecipeIngredientList> RecipeIngredientLists{get;set;}

    }
}